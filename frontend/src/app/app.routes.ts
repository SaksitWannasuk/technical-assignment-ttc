import { Routes } from '@angular/router';
import { PersonList } from './pages/person-list/person-list';

export const routes: Routes = [
  {
    path: '',
    redirectTo: 'person-list',
    pathMatch: 'full',
  },
  {
    path: 'person-list',
    component: PersonList
  }
];