# Simple POC for possible implementation of Beam tour with react/typescript

To run both Beam Playground frontend and react Tour of Beam locally in docker


Build image for playground frontend

```
cd beam/playground/frontend
docker build -t oborisevich/playground_frontend:latest .
```

Build image for react tour of beam POC

```
cd beam/playground/react-tour-poc
docker build -t oborisevich/react_tour_poc:latest .
```

Start both containers

```
cd beam/playground/react-tour-poc
docker-compose up
```

Stop both containers

```
cd beam/playground/react-tour-poc
docker-compose down
```

Go to localhost:3002 for Beam Playground Frontend
Go to localhost:3001 for React Tour POC with embedded Beam Playground




