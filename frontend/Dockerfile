FROM node:22-alpine
 
# Set working directory
WORKDIR /app
 
# Install dependencies first (better caching)
COPY package*.json ./
RUN npm install
 
# Copy the rest of the source
COPY . .
 
# Expose Vite's default dev port
EXPOSE 5173
 
# Run Vite dev server, accessible outside the container
CMD ["npm", "run", "dev", "--", "--host"]
