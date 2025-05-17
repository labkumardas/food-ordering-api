# Food Ordering API

This project implements the Food Ordering API based on the OpenAPI 3.1 specification.  
It supports product listing, order placement with promo code validation, and follows the API contract closely.

---

## Features

- List products  
- Get product details by ID  
- Place orders with optional promo codes  
- Validate promo codes from local coupon files  
- API key-based authentication

---

## Getting Started

### Prerequisites

- Go 1.18+ installed on your system  
- `git` installed  
- Download the promo code files:  
  - `couponbase1.gz`  
  - `couponbase2.gz`  
  - `couponbase3.gz`  

### Setup

1. Clone the repository:

   ```bash
   git clone git@github.com:labkumardas/food-ordering-api.git
   https://github.com/labkumardas/food-ordering-api.git
  


### Note on Coupon Data Loading

  The coupon data files (couponbase1.gz, couponbase2.gz, couponbase3.gz) are large (around 1 GB total) and require significant time to load and initialize. This delay occurs because:

  The files are compressed and need to be decompressed and parsed.

  Coupon data is loaded into memory structures used for fast validation.

  Loading happens asynchronously in the background to avoid blocking the API server startup.

### Important:
  Until the coupon data is fully loaded, promo code validation will not be available, and orders using promo codes may fail with the message:
  "coupon data not ready, try again later".

  To improve user experience, clients should handle this scenario by retrying later or placing orders without promo codes until the system is ready.
