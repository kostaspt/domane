import http from '@store/http'

export function search(query) {
  return http.get(`search/extensions?query=${query}`);
}
