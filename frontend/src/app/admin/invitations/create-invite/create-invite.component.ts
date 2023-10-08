import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
    selector: 'app-create-invite',
    standalone: true,
    imports: [CommonModule, FormsModule, ReactiveFormsModule],
    templateUrl: './create-invite.component.html'
})
export class CreateInviteComponent {
    role = '';
    email = '';
    subject = '';
    message = '';

    constructor(private router: Router) {}

    onSubmit(): void {
        /*const formData = { 
      role: this.role,
      email: this.email,
      subject: this.subject,
      message: this.message  
    };
    const jsonData = JSON.stringify(formData);*/
        //console.log(jsonData);
        this.router.navigate(['/']);
    }
}
