import { Component, EventEmitter, OnInit, Output } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SearchDTO } from '../../models/drive-dto';

@Component({
  selector: 'app-search-form',
  templateUrl: './search-form.component.html',
  styleUrls: ['./search-form.component.scss']
})
export class SearchFormComponent implements OnInit {

  form: FormGroup;

  @Output()
  searchEvent = new EventEmitter<SearchDTO>();

  constructor(private fb: FormBuilder) {
    this.form = this.fb.group({
      departure: [null, Validators.required],
      destination: [null, Validators.required],
    });
  }

  submit() {  
    const searchDTO: SearchDTO = {
      departure_location: this.form.value.departure,
      destination: this.form.value.destination,
      page: -1,
      size: -1
    } ;

    this.searchEvent.emit(searchDTO);
  }

  ngOnInit(): void {}

}
