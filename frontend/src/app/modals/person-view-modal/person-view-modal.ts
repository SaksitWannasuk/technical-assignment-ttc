import { Component, Inject } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';

import { MAT_DIALOG_DATA, MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatButtonModule } from '@angular/material/button';
import { Person } from '../../services/person';

@Component({
  selector: 'app-person-view-modal',
  imports: [
    CommonModule,
    DatePipe,
    MatDialogModule,
    MatButtonModule
  ],
  templateUrl: './person-view-modal.html',
  styleUrl: './person-view-modal.scss'
})
export class PersonViewModal {
  constructor(
    @Inject(MAT_DIALOG_DATA) public person: Person,
    private dialogRef: MatDialogRef<PersonViewModal>
  ) { }

  close(): void {
    this.dialogRef.close();
  }
}