# Xray-core

Forked version with some improvements. More info in original [repo](https://github.com/XTLS/Xray-core).

## Improvements

### Limit device count per client using local policy.

```json lines
{
  // ...
  "inbounds": [
    {
      // ...
      "settings": {
        // ...
        "clients": [
          {
            "id": "Your UUID",
            "email": "test@example.com", // <- required for stats
            "level": 0,
          }
        ]
      }
    }
  ],
  // ...
  "policy": {
    "levels": {
      "0": {
        "statsUserOnline": true, // <- required for device limitation
        "deviceCount": 1 // 0 or less for unlimited device count
      }
    }
  },
  "stats": {}, // <- required for enable stats
  "limiter": {} // <- required for enable limits
}
```

## Donation

- **ETH: `0xbdD857a41507006F648Cb08DCaf57b36Cb1560c8`**
- **BTC: `bc1q94jqe2gvm8zjtfc4h0m4p6jt3dlkasjjxrhaph`**
- **TRX: `TMzXAmfxQRYyNHMaCQWqckkY6Y97oqrhWf`**

## License

[Mozilla Public License Version 2.0](https://github.com/XTLS/Xray-core/blob/main/LICENSE)