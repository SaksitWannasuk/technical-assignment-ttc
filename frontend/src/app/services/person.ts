import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface Person {
  id: number;
  firstName: string;
  lastName: string;
  birthDate: string;
  age: number;
  address: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreatePersonRequest {
  firstName: string;
  lastName: string;
  birthDate: string;
  address: string;
}

@Injectable({
  providedIn: 'root'
})
export class PersonService {
  private apiUrl = 'http://localhost:3000/api';

  constructor(private http: HttpClient) {}

  public getPeople(): Observable<Person[]> {
    return this.http.get<Person[]>(`${this.apiUrl}/people`);
  }

  public getPersonById(id: number): Observable<Person> {
    return this.http.get<Person>(`${this.apiUrl}/person/${id}`);
  }

  public createPerson(data: CreatePersonRequest): Observable<Person> {
    return this.http.post<Person>(`${this.apiUrl}/person`, data);
  }
}