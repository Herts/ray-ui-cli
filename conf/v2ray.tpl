{
        "stats": {},
        "api": {
                "tag": "api",
                "services": [
                        "StatsService",
                        "HandlerService"
                ]
        },
        "policy": {
                "levels": {
                        "0": {
                                "statsUserUplink": true,
                                "statsUserDownlink": true
                        }
                },
                "system": {
                        "statsInboundUplink": true,
                        "statsInboundDownlink": true
                }
        },
        "log": {
                "loglevel": "warning",
                "access": "/var/log/v2ray/access.log",
                "error": "/var/log/v2ray/error.log"
        },
        "inbounds": [
                {
                        "tag": "proxy",
                        "port": 32516,
                        "protocol": "vmess",
                        "settings": {
                                "clients": [
                                ]
                        },
                        "streamSettings": {
                                "network": "ws"
                        }
                },
                {
                        "tag": "api",
                        "port": 8144,
                        "protocol": "dokodemo-door",
                        "settings": {
                                "address": "127.0.0.1"
                        }
                }
        ],
        "outbounds": [
                {
                        "protocol": "freedom",
                        "settings": {}
                }
        ],
        "routing": {
                "rules": [
                        {
                                "type": "field",
                                "inboundTag": [
                                        "api"
                                ],
                                "outboundTag": "api"
                        }
                ]
        }
}