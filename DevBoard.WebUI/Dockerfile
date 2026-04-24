# Build stage
FROM node:20.20.2-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci --legacy-peer-deps

COPY . .

# Runtime stage
FROM node:20.20.2-alpine

WORKDIR /app

# Create non-root user
RUN addgroup -g 1001 -S nodejs && adduser -S nodejs -u 1001

# Copy only production dependencies from builder
COPY --from=builder --chown=nodejs:nodejs /app/node_modules ./node_modules
COPY --from=builder --chown=nodejs:nodejs /app/package*.json ./
COPY --chown=nodejs:nodejs . .

ENV HOST=0.0.0.0
ENV PORT=5173

EXPOSE 5173
 
# Switch to non-root user
USER nodejs

CMD ["sh", "-c", "npm run dev -- --host ${HOST} --port ${PORT}"]
