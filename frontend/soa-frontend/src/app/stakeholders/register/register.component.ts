import { ChangeDetectionStrategy, Component } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { UserService } from '../user.service';
import { User, UserRole } from '../model/User.model';
import { Router } from '@angular/router';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css'],
})
export class RegisterComponent{
  roles: UserRole[] = ['Tourist', 'Guide'];

  registerForm = this.fb.group({
    username: ['', Validators.required],
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required, Validators.minLength(6)]],
    role: ['Tourist', Validators.required] // default role
  });

  constructor(private fb: FormBuilder, private userService: UserService, private router: Router) {}

  onSubmit() {
    if (this.registerForm.valid) {
      const newUser: User = this.registerForm.value as User;
      this.userService.register(newUser).subscribe({
        next: () => {
          alert('Registration successful!');
          this.router.navigate(['/']);
        },
        error: err => alert('Registration failed: ' + err.error)
      });
    }
  }

}
