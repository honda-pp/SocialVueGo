import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';

const userApi = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

export const loginUser = async (userData) => {
  try {
    const response = await userApi.post(`/login`, userData);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const logoutUser = async () => {
  try {
    const response = await userApi.post(`/logout`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const checkLoggedIn = async () => {
  try {
    const response = await userApi.get(`/checkLoggedIn`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const signupUser = async () => {
  try {
    const response = await userApi.get(`/signup`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};
