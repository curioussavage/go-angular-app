import { Component, Input, Output, EventEmitter } from '@angular/core';
import { ModelsUser } from 'projects/api-client';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialogComponent } from '../confirm-dialog/confirm-dialog.component';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.css'],
})
export class UserCardComponent {
  @Input() user!: ModelsUser;
  @Output() delete = new EventEmitter<number>();

  constructor(public dialog: MatDialog) { }

  openDeleteDialog() {
    const dialogRef = this.dialog.open(ConfirmDialogComponent, {
      data: { message: 'Are you sure you want to delete?' }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.delete.emit(this.user.userID);
      }
    });

  }
}
