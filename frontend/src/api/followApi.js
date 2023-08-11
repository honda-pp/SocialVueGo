import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';

const userApi = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});

export const followUser = async (userId) => {
  try {
    const response = await userApi.post(`/follow/${userId}`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const unfollowUser = async (userId) => {
  try {
    const response = await userApi.post(`/unfollow/${userId}`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const getFollowingIDs = async () => {
  try {
    const response = await userApi.get(`/followingIDs`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};

export const getFollowerIDs = async () => {
  try {
    const response = await userApi.get(`/followerIDs`);
    return response.data;
  } catch (error) {
    throw error.response.data;
  }
};
