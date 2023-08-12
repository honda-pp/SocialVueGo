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

export const getSessionInfo = async () => {
  try {
    const response = await userApi.get(`/sessionInfo`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const getUserInfo = async (userID) => {
  try {
    const response = await userApi.get(`/userInfo/${userID}`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const signupUser = async (userData) => {
  try {
    const response = await userApi.post(`/signup`, userData);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const getUserList = async (userID, listType) => {
  let path = ``
  if (listType == null) {
    path = `/userList`
  } else if (listType === 'Follower') {
    path = `/followerUserList/${userID}`
  } else {
    path = `/followingUserList/${userID}`
  }
  try {
    const response = await userApi.get(path);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};
