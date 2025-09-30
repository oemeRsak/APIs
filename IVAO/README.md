# API Template

## Description

*Provide a short summary of what the API does.*

Example:
This API allows clients to fetch information about **\[resource]**.

---

## Base URL

The base URL for all API requests is:

```url
https://your-api-domain.com
```

---

## Endpoints

### `METHOD /endpoint`

**Description**:
*Short explanation of what this endpoint does.*

#### Parameters

* `param1` (required/optional): *Description of the parameter.*
* `param2` (required/optional): *Description of the parameter.*

#### Response

Returns a JSON object with the following properties:

* `property1`: *Description of the property.*
* `property2`: *Description of the property.*

#### Example

**Request:**

```http
METHOD /endpoint?param1=value&param2=value
```

**Response:**

```json
{
  "property1": "example",
  "property2": "example"
}
```

---

## Errors

This API uses the following error codes:

* `400 Bad Request`: The request was malformed or missing required parameters.
* `401 Unauthorized`: The API key provided was invalid or missing.
* `403 Forbidden`: The client does not have access rights to the resource.
* `404 Not Found`: The requested resource was not found.
* `500 Internal Server Error`: An unexpected error occurred on the server.

---

Would you like me to expand this into a **multi-endpoint template** (with placeholders for GET, POST, PUT, DELETE) so you can just fill in details per endpoint?
