import { Component, Input, OnInit } from '@angular/core';
import { CarDTO } from 'src/modules/shared/models/car-dto';

@Component({
  selector: 'app-car-viewer',
  templateUrl: './car-viewer.component.html',
  styleUrls: ['./car-viewer.component.scss']
})
export class CarViewerComponent implements OnInit {

  @Input()
  car: CarDTO;

  constructor() {
    this.car = { 
      plateNumber:"",
      brand:"",
      carModel:"",
      fuelConsumption: 0.0,
      volume: 0.0,
      power: 0.0,
    }
   }

  ngOnInit(): void {}

}
