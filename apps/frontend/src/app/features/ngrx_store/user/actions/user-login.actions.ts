import { createActionGroup, emptyProps, props } from '@ngrx/store';
import { UserModel } from 'shared/models/user.model';

export const userLoginActions = createActionGroup({
  source: 'Login page',
  events: {
    loggedIn: emptyProps(),
  },
});
