import { Injectable } from '@angular/core';
import {
  ActivatedRouteSnapshot,
  CanDeactivate,
  GuardResult,
  MaybeAsync,
  Router,
  RouterStateSnapshot,
} from '@angular/router';
import { ConfirmationService, MessageService, PrimeIcons } from 'primeng/api';
import { Subject } from 'rxjs';
import { CUstomDialogsKeys, CustomToastsKeys } from 'shared/components';
import { SeverityTypes } from 'shared/components/custom_toasts/enums/severity-types.enum';
import { IFormEmptyDataSafe } from 'shared/interfaces';

@Injectable({
  providedIn: 'root',
})
export class SaveFormDataGuard implements CanDeactivate<IFormEmptyDataSafe> {
  constructor(
    private confirmationService: ConfirmationService,
    private messageService: MessageService,
    private router: Router // Inject Router
  ) {}

  confirmationSubject = new Subject<boolean>();

  canDeactivate(
    component: IFormEmptyDataSafe,
    currentRoute: ActivatedRouteSnapshot,
    currentState: RouterStateSnapshot,
    nextState: RouterStateSnapshot // Next state is the target route
  ): MaybeAsync<GuardResult> {
    const targetRoute = nextState.url;
    if (
      (targetRoute.includes('login') || targetRoute.includes('register')) &&
      !component.isFormEmpty()
    ) {
      this.confirmationService.confirm({
        header: 'Confirmation',
        acceptIcon: PrimeIcons.CHECK,
        rejectIcon: PrimeIcons.TIMES,
        icon: PrimeIcons.INFO_CIRCLE,
        key: CUstomDialogsKeys.BasicConfirmDialog,
        message: 'Leaving will discard the values entered in the form',
        accept: () => {
          // If user clicks "Yes"
          this.messageService.add({
            key: CustomToastsKeys.BasicToast,
            severity: SeverityTypes.INFO,
            summary: 'Confirmed',
            detail: 'You have accepted discarding the data',
          });
          this.confirmationSubject.next(true);
          this.confirmationSubject.complete();
        },
        reject: () => {
          // If user clicks "No"
          this.messageService.add({
            key: CustomToastsKeys.BasicToast,
            severity: SeverityTypes.WARN,
            summary: 'Cancelled',
            detail: 'You have rejected discarding the data',
          });
          this.confirmationSubject.next(false);
          this.confirmationSubject.complete();
        },
      });

      return this.confirmationSubject.asObservable();
    }
    return true;
  }
}
