# Heroku Log Visualizer

A **Heroku log visualization tool** that provides a real-time dashboard for monitoring logs efficiently.

---

## Features

-  **Receive logs via HTTP endpoint** (Heroku Log Drains)
-  **Store logs** in PostgreSQL for persistence
-  **Filter logs** by level (INFO, WARNING, ERROR,)
-  **Real-time updates** via WebSockets (future)
-  **Easily deployable** via Docker

---

## Project Structure
heroku-log-visualizer/ │── backend/ # FastAPI backend │ ├── app.py # Main API logic │ ├── database.py # Database connection │ ├── log_processor.py # Handles log parsing/filtering │ ├── requirements.txt # Python dependencies │ ├── Dockerfile # Backend containerization │── frontend/ # React/Vue dashboard │ ├── src/ │ ├── public/ │ ├── package.json # Frontend dependencies │ ├── Dockerfile # Frontend containerization │── .env.example # Environment variables template │── docker-compose.yml # One-command setup │── README.md # This file


---

##  Future Enhancements

- [ ] **WebSockets for real-time logs** instead of polling  
- [ ] **Log Search & Advanced Filtering** (by app, level, timeframe, keyword)  
- [ ] **GraphQL API Support** for better query performance  
- [ ] **AI-based Log Analysis** (Anomaly detection, error prediction)  
- [ ] **Deploy a hosted version**  
- [ ] **CLI Tool**  

---


