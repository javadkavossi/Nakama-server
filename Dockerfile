FROM heroiclabs/nakama:3.25.0

WORKDIR /nakama/data/modules
COPY backend.so .
COPY local.yml /nakama/data/
COPY *.json /nakama/data/modules


# GOARCH=arm64 GOOS=darwin go build -trimpath -buildmode=plugin -o backend.so
# docker-compose up --build


# go build --trimpath --buildmode=plugin -o backend.so  
