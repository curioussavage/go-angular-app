import { Component, Input } from '@angular/core';
import { User } from '../user';
import { ModelsUser } from 'projects/api-client';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.css'],
})
export class UserCardComponent {
  @Input() user!: ModelsUser;

}
