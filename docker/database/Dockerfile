# Menggunakan image resmi PostgreSQL sebagai base image
FROM postgres:13-alpine

# Set environment variables untuk PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=postgres

# Menyalin file konfigurasi atau skrip SQL jika diperlukan
# Misalnya, Anda bisa menambahkan file SQL untuk menginisialisasi database
# COPY init.sql /docker-entrypoint-initdb.d/

# Menentukan port yang akan digunakan
EXPOSE 5434

# Perintah untuk menjalankan PostgreSQL
CMD ["postgres"]
