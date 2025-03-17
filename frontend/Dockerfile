# Use Node.js official image to build the React app
FROM node:20 as builder

# Set the working directory
WORKDIR /app

# Copy package files first to install dependencies
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire project to the container
COPY . .

# Build the React app
RUN npm run build

# Use Nginx to serve the React build
FROM nginx:1.25

# Copy the built React app to Nginx's public directory
COPY --from=builder /app/build /usr/share/nginx/html

# Expose port 80 (default Nginx port)
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
