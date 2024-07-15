# Use a Node.js base image
FROM node:16

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code to the working directory
COPY . .

# Build the Quasar application
RUN npm run build

# Install http-server to serve the built files
RUN npm install -g http-server

# Expose port 8080
EXPOSE 8080

# Serve the built files using http-server
CMD ["http-server", "dist/spa", "-p", "8080"]
