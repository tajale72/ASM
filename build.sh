# touch index.html index.js index.css main.go Dockerfile 

# Build the Docker image
docker build -t asm .

# Run the Docker container
docker run -p 8080:8080 asm



