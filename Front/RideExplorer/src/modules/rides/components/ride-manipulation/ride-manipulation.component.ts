import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { UtilService } from 'src/modules/shared/services/util.service';
import { DriveDTO, NewDriveDTO, UpdateDriveDTO } from '../../models/drive-dto';

@Component({
  selector: 'app-ride-manipulation',
  templateUrl: './ride-manipulation.component.html',
  styleUrls: ['./ride-manipulation.component.scss']
})
export class RideManipulationComponent implements OnInit {

  @Input()
  mode: string;

  @Input()
  ride: DriveDTO | undefined;

  form: FormGroup;

  @Output()
  submitCreate = new EventEmitter<NewDriveDTO>();

  @Output()
  submitUpdate = new EventEmitter<UpdateDriveDTO>();

  constructor(private fb: FormBuilder, private utilService: UtilService) {
    this.mode = "create";
    this.form = this.fb.group({
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

    if(this.ride)
      this.form.setValue(this.ride);

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
      departure_date_time: this.form.value.departure_date_time,
      departure_address: this.form.value.departure_address,
      free_places: this.form.value.free_places,
      planned_arrival_time: this.form.value.planned_arrival_time,
      note: this.form.value.note,
      distance: this.form.value.distance,
    };

    this.submitCreate.emit(dto);
  }

  updateDrive() {
    if(!this.ride) return;
    
    const dto: UpdateDriveDTO = {
      id: this.ride?.id,
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

}
