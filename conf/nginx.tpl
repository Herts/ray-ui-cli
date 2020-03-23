server {
    listen {{.Port}};
    server_name {{.ServerName}}; #edit here
    location /{{.Mask}} {
        proxy_pass       http://localhost:{{.RayPort}}; #edit here
        proxy_redirect   off;
        proxy_set_header Host $host;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}