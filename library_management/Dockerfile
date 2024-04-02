# PostgreSQL için resmi Docker imajını kullanın
FROM postgres:latest

# PostgreSQL ayarları ve veritabanı oluşturma işlemleri burada yapılabilir

# Go uygulamasını çalıştırmak için gerekli olan ek bağımlılıkları yükleyin
RUN apt-get update && apt-get install -y \
    golang-go \
    git \
    && rm -rf /var/lib/apt/lists/*

# Go uygulamasının çalıştırılacağı dizini oluşturun
WORKDIR /go/src/app

# Go uygulamanızı konteynera kopyalayın
COPY . .

# Go bağımlılıklarını yükleyin


# Go uygulamanızı derleyin
RUN go build -o main .

# Docker konteynerı çalıştırılırken uygulamanın otomatik olarak başlaması için bir komut belirtin
CMD ["./main"]
