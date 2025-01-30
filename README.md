# Financial Insights Web Application

## Overview
This project is a web application designed to intelligently read, analyze, and organize banking data for enhanced financial insights. It leverages AI-powered processing with a **Go backend**, a **React frontend**, and various AI services to clean and structure financial data efficiently.

## Architecture
The application follows a **microservices** architecture consisting of the following components:

1. **AI Wrapper Service (Go)**: A microservice that integrates with AI providers like OpenAI and DeepSeek for data processing and intelligent analysis.
2. **API Service (Go)**: The primary backend service that handles business logic, controllers, and interactions with the database.
3. **Frontend (React)**: A web interface to visualize charts, insights, and banking data analysis.
4. **PostgreSQL Database**: A relational database to store cleaned and structured financial data.
5. **Redis**: A caching layer to optimize performance and handle real-time data efficiently.
6. **Kafka**: A message broker to facilitate communication between microservices asynchronously.
7. **Docker & Docker Compose**: The project is containerized for easy deployment and scalability.

## Tech Stack
- **Backend:** Golang (Go)
- **Frontend:** React.js
- **Database:** PostgreSQL
- **Cache:** Redis
- **Message Broker:** Kafka
- **AI Services:** OpenAI, DeepSeek, and others
- **Containerization:** Docker, Docker Compose

## Features
- **User Banking Data Integration**: Collects financial data from users securely.
- **AI-Powered Cleanup & Organization**: Uses AI services to clean and structure the data.
- **Financial Insights & Analytics**: Provides users with enhanced financial insights through interactive charts and reports.
- **Scalable & Modular Architecture**: Microservices-based design for improved scalability and maintainability.
- **Efficient Communication with Kafka**: Ensures smooth and asynchronous data processing.
- **Containerized Deployment**: Uses Docker and Docker Compose for easy setup and deployment.

## Setup & Installation
### Prerequisites
Ensure you have the following installed:
- Docker & Docker Compose
- Go (latest version)
- Node.js & npm/yarn
- Kafka & Redis
- PostgreSQL

### Installation Steps
1. **Clone the repository**
   ```sh
   git clone https://github.com/your-repo/financial-insights.git
   cd financial-insights
   ```

2. **Start services using Docker Compose**
   ```sh
   docker-compose up -d
   ```

3. **Run Backend Services**
   - Start the AI Wrapper Service:
     ```sh
     cd services/ai-wrapper
     go run main.go
     ```
   - Start the API Service:
     ```sh
     cd services/api
     go run main.go
     ```

4. **Run Frontend**
   ```sh
   cd frontend
   npm install  # or yarn install
   npm start  # or yarn start
   ```

5. **Access the Application**
   Open your browser and navigate to `http://localhost:3000` to use the application.

## API Endpoints
### AI Wrapper Service
- `POST /process-data` - Sends banking data for AI-based cleaning and structuring.

### API Service
- `GET /transactions` - Fetches user transactions.
- `POST /upload` - Uploads new banking data.
- `GET /insights` - Retrieves AI-generated financial insights.

## Kafka Topics
- `banking-data-raw` - Raw banking data before processing.
- `banking-data-cleaned` - AI-processed and structured financial data.
- `insights-generated` - AI-generated insights for visualization.

## Environment Variables
Ensure you have a `.env` file in the root directory with the following variables:
```env
DATABASE_URL=postgres://user:password@db:5432/financial_db
REDIS_URL=redis://redis:6379
KAFKA_BROKER=kafka:9092
OPENAI_API_KEY=your_openai_key
DEEPSEEK_API_KEY=your_deepseek_key
```

## Contributing
1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -m 'Add new feature'`).
4. Push to your branch (`git push origin feature-name`).
5. Open a Pull Request.

## License
This project is licensed under the MIT License.


