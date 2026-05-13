import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';

import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { Person, PersonService } from '../../services/person';

import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { PersonModal } from '../../modals/person-modal/person-modal';
import { PersonViewModal } from '../../modals/person-view-modal/person-view-modal';


@Component({
  selector: 'app-person-list',
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatDialogModule
  ],
  templateUrl: './person-list.html',
  styleUrl: './person-list.scss'
})
export class PersonList implements OnInit {
  people: Person[] = [];

  displayedColumns: string[] = [
    'id',
    'fullName',
    'address',
    'birthDate',
    'age',
    'actions'
  ];

  constructor(private personService: PersonService, private dialog: MatDialog) { }

  ngOnInit(): void {
    this.loadPeople();
  }

  public loadPeople(): void {
    this.personService.getPeople().subscribe({
      next: (data) => {
        this.people = data;
      },
      error: (err) => {
        console.error('Failed to load people', err);
      }
    });
  }

  public openAddModal(): void {
    const dialogRef = this.dialog.open(PersonModal, {
      width: '500px'
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result === true) {
        this.loadPeople();
      }
    });
  }

  public openViewModal(person: Person): void {
  this.dialog.open(PersonViewModal, {
    width: '500px',
    data: person
  });
}
}