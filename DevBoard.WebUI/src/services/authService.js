const API_BASE_URL = 'http://localhost:8080/api/v1';

function formatValidationDetails(details) {
  if (!Array.isArray(details) || details.length === 0) return '';

  return details
    .map((item) => {
      const field = item?.field || 'Field';
      const message = item?.message || 'Invalid value';

      if (field === 'Password' && message.includes("min")) {
        return 'Password must be at least 8 characters.';
      }

      if (field === 'Email') {
        return 'Please enter a valid email.';
      }

      return `${field}: ${message}`;
    })
    .join(' ');
}

function extractErrorMessage(data) {
  if (!data) return 'Something went wrong.';

  
  if (data.error && typeof data.error === 'object') {
    const detailsMessage = formatValidationDetails(data.error.details);
    if (detailsMessage) return detailsMessage;

    if (data.error.code === 'VALIDATION_FAILED') {
      return data.error.message || 'Validation failed.';
    }

    if (data.error.code === 'USER_ALREADY_EXISTS') {
      return 'An account with this email already exists.';
    }

    if (typeof data.error.message === 'string') {
      return data.error.message;
    }
  }

  if (Array.isArray(data.details) && data.details.length > 0) {
    const detailsMessage = formatValidationDetails(data.details);
    if (detailsMessage) return detailsMessage;
  }

  if (typeof data.message === 'string') return data.message;
  if (typeof data.error === 'string') return data.error;
  if (typeof data.detail === 'string') return data.detail;

  return 'Something went wrong.';
}

async function parseResponse(response) {
  const data = await response.json().catch(() => ({}));

  if (!response.ok) {
    throw new Error(extractErrorMessage(data));
  }

  return data;
}

export async function loginUser(payload) {
  const response = await fetch(`${API_BASE_URL}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(payload),
  });

  return parseResponse(response);
}

export async function signupUser(payload) {
  const signupPayload = {
    firstname: payload.firstName,
    lastname: payload.lastName,
    email: payload.email,
    password: payload.password,
    phoneNumber: payload.phoneNumber || '',
    role: payload.role,
  };

  const response = await fetch(`${API_BASE_URL}/auth/signup`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include',
    body: JSON.stringify(signupPayload),
  });

  return parseResponse(response);
}
export async function fetchRoles() {
  return [
    { id: 1, name: 'Developer' },
    { id: 2, name: 'Recruiter' },
  ];
}