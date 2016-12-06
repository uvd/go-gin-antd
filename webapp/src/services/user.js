import request from '../utils/request';

export async function query() {
  return request('/api/users');
}

export async function list() {
  return request('/api/user/list');
}
export async function login() {
  return request('/api/user/login');
}

export async function regist() {
  return request('/api/user/regist');
}


