import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators
} from '@angular/forms';

import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';

import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';

import { MatButtonModule } from '@angular/material/button';

import { MatDatepickerModule } from '@angular/material/datepicker';
import { MatNativeDateModule } from '@angular/material/core';
import { PersonService } from '../../services/person';


@Component({
  selector: 'app-person-modal',
  imports: [
    CommonModule,
    ReactiveFormsModule,

    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,

    MatDatepickerModule,
    MatNativeDateModule
  ],
  templateUrl: './person-modal.html',
  styleUrl: './person-modal.scss'
})
export class PersonModal {
  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private personService: PersonService,
    private dialogRef: MatDialogRef<PersonModal>
  ) {
    this.form = this.fb.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      birthDate: ['', Validators.required],
      address: ['', Validators.required]
    });
  }

  public submit(): void {
    if (this.form.invalid) {
      return;
    }

    const formValue = this.form.value;

    const payload = {
      firstName: formValue.firstName,
      lastName: formValue.lastName,
      birthDate: this.formatDate(formValue.birthDate),
      address: formValue.address
    };

    this.personService.createPerson(payload).subscribe({
      next: () => {
        this.dialogRef.close(true);
      },
      error: (err) => {
        console.error('Failed to create person', err);
      }
    });
  }

  public close(): void {
    this.dialogRef.close(false);
  }

  public formatDate(date: Date): string {
    return date.toISOString().split('T')[0];
  }
}