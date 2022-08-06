import { AfterViewInit, Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { AuthService } from 'src/modules/auth/services/auth-service/auth.service';
import { SnackBarService } from 'src/modules/shared/services/snack-bar.service';

@Component({
  selector: 'app-header-common',
  templateUrl: './header-common.component.html',
  styleUrls: ['./header-common.component.scss']
})
export class HeaderCommonComponent implements AfterViewInit {

  role: string;

  constructor(private authService: AuthService) {
    this.role = '';
  }

  ngAfterViewInit() {}

  logout() {
    this.authService.logout();
  }

}
