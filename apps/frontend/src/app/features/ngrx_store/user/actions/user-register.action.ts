import { createActionGroup, emptyProps, props } from '@ngrx/store';
import { UserModel } from 'shared/models/user.model';

export const userRegisterActions = createActionGroup({
  source: 'Register page',
  events: {
    loggedIn: emptyProps(),
  },
});
