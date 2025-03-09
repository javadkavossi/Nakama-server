# FROM heroiclabs/nakama:3.25.0

# # ایجاد ساختار دایرکتوری مورد نیاز
# RUN mkdir -p /nakama/data/modules/data/modules

# # کپی ماژول در مسیر صحیح
# COPY backend.so /nakama/data/modules/data/modules/
# COPY local.yml /nakama/data/
# COPY *.json /nakama/data/modules


# ----------------------------

# GOARCH=arm64 GOOS=darwin go build -trimpath -buildmode=plugin -o backend.so
# docker-compose up --build



# go build --trimpath --buildmode=plugin -o backend.so  



FROM heroiclabs/nakama:3.25.0

# Create the correct directory structure
RUN mkdir -p /nakama/data/modules/data/modules

# Copy files to the correct locations
COPY backend.so /nakama/data/modules/data/modules/
COPY local.yml /nakama/data/
COPY *.json /nakama/data/modules/data/modules/