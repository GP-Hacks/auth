#!/bin/bash
# scripts/deploy.sh

set -e

ENVIRONMENT="$1"
IMAGE="$2"
CONTAINER_NAME="$3"
PORT="$4"
YC_OAUTH_TOKEN="$5"

PROJECT_NAME="kdt-auth-service"

if [[ -z "$ENVIRONMENT" || -z "$IMAGE" || -z "$CONTAINER_NAME" || -z "$PORT" || -z "$YC_OAUTH_TOKEN" ]]; then
    echo "❌ Error: Missing required arguments"
    echo "Usage: $0 <environment> <image> <container_name> <port> <yc_oauth_token>"
    exit 1
fi

DEPLOY_PATH="/opt/apps/$ENVIRONMENT"
COMPOSE_FILE="docker-compose.$ENVIRONMENT.yml"
NETWORK_NAME="${PROJECT_NAME}-network-$ENVIRONMENT"

echo "🚀 Starting deployment..."
echo "Environment: $ENVIRONMENT"
echo "Image: $IMAGE"
echo "Container: $CONTAINER_NAME"
echo "Port: $PORT"

echo "📁 Creating deployment directory..."
sudo mkdir -p "$DEPLOY_PATH"
cd "$DEPLOY_PATH"

if ! command -v yc &> /dev/null; then
    echo "📥 Installing Yandex Cloud CLI..."
    curl -sSL https://storage.yandexcloud.net/yandexcloud-yc/install.sh | bash
    export PATH="$HOME/yandex-cloud/bin:$PATH"
fi

echo "🔑 Getting secrets from Lockbox..."
yc config set token "$YC_OAUTH_TOKEN"
yc config set cloud-id b1grt0fvgql5big8hevj
yc config set folder-id b1gq39fmv588jocgh7to

echo "📝 Creating .env file..."
yc lockbox payload get "${PROJECT_NAME}-secrets-$ENVIRONMENT" --format json | \
    jq -r '.entries[] | "\(.key)=\(.text_value)"' | sudo tee .env > /dev/null

echo "🔗 Ensuring network exists..."
sudo docker network create "$NETWORK_NAME" 2>/dev/null || echo "Network exists"

echo "🐳 Creating docker compose file..."
cat << EOF | sudo tee "$COMPOSE_FILE" > /dev/null
services:
  app:
    image: $IMAGE
    container_name: $CONTAINER_NAME
    ports:
      - "$PORT:3000"
    env_file:
      - .env
    networks:
      - $NETWORK_NAME
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s

networks:
  $NETWORK_NAME:
    external: true
EOF

echo "🔑 Authenticating with Yandex Container Registry..."
echo "$YC_OAUTH_TOKEN" | sudo docker login \
  --username oauth \
  --password-stdin \
  cr.yandex

echo "📦 Pulling latest image..."
sudo docker pull "$IMAGE"

echo "🛑 Stopping old containers..."
sudo docker stop "$CONTAINER_NAME" 2>/dev/null || true
sudo docker rm "$CONTAINER_NAME" 2>/dev/null || true

echo "▶️  Starting new container..."
sudo docker compose -f "$COMPOSE_FILE" up -d

echo "🏥 Waiting for health check..."
for i in {1..10}; do
    if curl -f "http://localhost:$PORT/health" 2>/dev/null; then
        echo "✅ Application ready!"
        break
    fi
    
    if [ $i -eq 10 ]; then
        echo "❌ Health check failed"
        exit 1
    fi
    
    sleep 30
done

echo "✅ Deployment completed successfully!"
echo "🎉 Application is ready at http://localhost:$PORT"
