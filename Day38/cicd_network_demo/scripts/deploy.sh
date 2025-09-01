#!/bin/bash

# CI/CD Deployment Script
set -e

echo "🚀 Starting deployment process..."
echo "📦 Application: CI/CD Network Demo"
echo "🌐 Environment: production"

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Build the application
echo "🔨 Building application..."
docker build -t cicd-network-demo:latest .

# Run tests in CI mode
echo "🧪 Running tests..."
docker run --rm -e CI=true cicd-network-demo:latest ./cicd-network-demo loadtest

# Deploy to production
echo "🚀 Deploying to production..."
docker-compose -f docker-compose.prod.yml up -d

echo "✅ Deployment completed successfully!"
echo "📊 Health check: http://localhost:8080/health"