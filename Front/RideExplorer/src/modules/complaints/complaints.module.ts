import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ComplaintsRoutes } from './complaints.routes';
import { SharedModule } from '../shared/shared.module';
import { MatIconModule } from '@angular/material/icon';
import { RouterModule } from '@angular/router';
import { ComplaintsPageComponent } from './pages/complaints-page/complaints-page.component';
import { ComplaintCardComponent } from './components/complaint-card/complaint-card.component';



@NgModule({
  declarations: [
    ComplaintsPageComponent,
    ComplaintCardComponent
  ],
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule.forChild(ComplaintsRoutes),
    SharedModule,
    FormsModule,
    MatIconModule,
  ]
})
export class ComplaintsModule { }
