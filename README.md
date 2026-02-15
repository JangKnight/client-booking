# Client Booking Project

## API Overview


| Method | Path | Summary | Auth | Role |
| --- | --- | --- | --- | --- |
| POST | /users | Sign up | Public | — |
| POST | /sessions | Log in | Public | — |
| GET | /me | Get current user | Required | Any |
| GET | /sellers | List sellers | Public | — |
| GET | /sellers/{sellerId} | Seller profile | Public | — |
| GET | /sellers/{sellerId}/services | List services for seller | Public | — |
| POST | /sellers/{sellerId}/services | Create service | Required | Seller |
| GET | /services/{serviceId} | Service details | Public | — |
| PATCH | /services/{serviceId} | Update service | Required | Seller |
| GET | /services/{serviceId}/slots | List availability slots | Public | — |
| POST | /services/{serviceId}/slots | Create availability slot | Required | Seller |
| DELETE | /slots/{slotId} | Remove availability slot | Required | Seller |
| POST | /bookings | Create booking for a slot | Required | Consumer |
| GET | /me/bookings | List my bookings | Required | Consumer |
| GET | /bookings/{bookingId} | Booking details | Required | Consumer |
| DELETE | /bookings/{bookingId} | Cancel booking | Required | Consumer |


