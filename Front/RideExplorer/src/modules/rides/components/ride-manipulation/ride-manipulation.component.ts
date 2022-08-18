import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO, NewDriveDTO, UpdateDriveDTO } from '../../models/drive-dto';
import * as moment from 'moment';

@Component({
  selector: 'app-ride-manipulation',
  templateUrl: './ride-manipulation.component.html',
  styleUrls: ['./ride-manipulation.component.scss']
})
export class RideManipulationComponent implements OnInit {

  @Input()
  mode: string;

  private _ride: DriveDTO | undefined;

  @Input() set ride(value: DriveDTO | undefined) {
    this._ride = value;

    if(this._ride)
      this.form.setValue(this._ride);

    const ddt = this.form.get("departure_date_time")?.value; 
    this.form.get("departure_date_time")?.setValue(this.formatNumberToDate(ddt));

    const pat = this.form.get("planned_arrival_time")?.value; 
    this.form.get("planned_arrival_time")?.setValue(this.formatNumberToDate(pat));
  }
 
  get ride(): DriveDTO | undefined {
      return this._ride;
  }

  form: FormGroup;

  @Output()
  submitCreate = new EventEmitter<NewDriveDTO>();

  @Output()
  submitUpdate = new EventEmitter<UpdateDriveDTO>();

  @Output()
  finishRideEvent = new EventEmitter<number>();

  dateValue: Date = new Date ("16/05/2017 13:00");

  constructor(private fb: FormBuilder, private utilService: UtilService) {
    this.mode = "create";
    this.form = this.fb.group({
      id: [null],
      finished: [null],
      driver_username: [null],
      departure_location: [null, Validators.required],
      destination: [null, Validators.required],
      departure_date_time: [null, Validators.required],
      departure_address: [null, Validators.required],
      free_places: ['', [Validators.required, Validators.min(1), Validators.pattern('^(0|[1-9][0-9]*)$')]],
      planned_arrival_time: [null, Validators.required],
      note: [null],
      distance: ['', [Validators.required, Validators.min(1), Validators.pattern('^(0|[1-9][0-9]*)$')]]
    });
  }

  ngOnInit(): void {
    if(this.mode === "create") return;

    if(this.mode === "update") {
      this.disableUpdate();
    }

    else if(this.mode === "view") {
      this.disableView();
    }
  }

  submit() {

    if(this.mode==="create") {
      this.createDrive();
    } 
    else if(this.mode==="update") {
      this.updateDrive();
    }
    
  }

  createDrive() {
    const dto: NewDriveDTO = {
      driver_username: this.utilService.getLoggedUserUsername(),
      departure_location: this.form.value.departure_location,
      destination: this.form.value.destination,
      departure_date_time: this.formatDateToNumber(this.form.value.departure_date_time),
      departure_address: this.form.value.departure_address,
      free_places: this.form.value.free_places,
      planned_arrival_time: this.formatDateToNumber(this.form.value.planned_arrival_time),
      note: this.form.value.note,
      distance: this.form.value.distance,
    };

    this.submitCreate.emit(dto);
  }

  updateDrive() {
    if(!this._ride) return;
    
    const dto: UpdateDriveDTO = {
      id: this._ride?.id,
      departure_address: this.form.value.departure_address,
      free_places: this.form.value.free_places,
      note: this.form.value.note,
    };

    this.submitUpdate.emit(dto);
  }

  disableUpdate() {
    this.form.get("driver_username")?.disable(); 
    this.form.get("departure_location")?.disable(); 
    this.form.get("destination")?.disable(); 
    this.form.get("departure_date_time")?.disable(); 
    this.form.get("planned_arrival_time")?.disable(); 
    this.form.get("distance")?.disable(); 
  }

  disableView() {
    this.disableUpdate();
    this.form.get("departure_address")?.disable(); 
    this.form.get("free_places")?.disable(); 
    this.form.get("note")?.disable(); 
  }

  formatDateToNumber(date: Date): number {
    return moment(date).valueOf();
  }

  formatNumberToDate(millis: number): string {
    const strDate = moment(millis).format("yyyy-MM-DDTHH:mm"); 
    const date = new Date(strDate);
    const isoDate = date.toISOString().slice(0,16);
    return isoDate;
  }

  finishRide() {
    this.finishRideEvent.emit(this._ride?.id);
  }

}
