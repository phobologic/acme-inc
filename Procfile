web: 
  command: acme-inc server
  ports:
    - "80:80"
    - "6063:80":
        protocol: "http"
    - "8080:8080":
        protocol: "tcp"
worker:
  command: acme-inc worker
