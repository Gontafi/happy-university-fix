import './css/login.css';
import { sendDataToBackend } from './components/Button';

const Login = () => {
    const handleSubmit = (e) => {
        e.preventDefault();
        const email = e.target.email.value;
        const password = e.target.password.value;
      
        // Send the data to your Go back-end
        sendDataToBackend(email, password);
    };

    return (
        <div>
            <div class="form-wrapper"> <h1>Log in</h1>
                <form handleSubmit={handleSubmit} action="#">
                    <div class="input-field">
                        <label for="email">Email address:</label>
                        <input type="email" id="email" name="email" placeholder="Email address"></input>
                    </div>
                    <div class="input-field">
                        <label for="password">Password:</label>
                        <input type="password" id="password" name="password" placeholder="Password"></input>
                    </div>
                    <a href="/forgotpassword">Forgot password?</a>
                    <button type="submit">Log in</button>
                </form>
            </div>
        </div>
    );
}

export default Login;
