import env from './.env';

export const environment = {
  production: true,
  gateway: 'http://localhost:4200/frontend-gateway', //need to change localhost:4200 to the ip or domain name of the server running nginx
  // version: env.npm_package_version,
};
