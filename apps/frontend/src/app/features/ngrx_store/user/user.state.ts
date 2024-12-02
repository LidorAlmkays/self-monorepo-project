import { createFeature, createReducer, on } from '@ngrx/store';
import { userRegisterActions } from './actions/user-register.action';
import { userLoginActions } from './actions/user-login.actions';

export interface UserState {
  loggedIn: boolean;
}

const initialState: UserState = {
  loggedIn: false,
};

const reducer = createReducer(
  initialState,
  on(userRegisterActions.loggedIn, (state) => ({
    ...state,
    loggedIn: true,
  })),
  on(userLoginActions.loggedIn, (state) => ({
    ...state,
    loggedIn: true,
  }))
);

export const userFeature = createFeature({
  name: 'userFeature',
  reducer,
  extraSelectors: ({}) => ({}),
});
