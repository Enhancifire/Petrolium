import 'package:firebase_auth/firebase_auth.dart';
import 'package:google_sign_in/google_sign_in.dart';

class AuthService {
  final FirebaseAuth _auth = FirebaseAuth.instance;
  final GoogleSignIn _googleSignIn = GoogleSignIn();

  Future<void> signInWithEmailAndPassword(
    String email,
    String password,
  ) async {}

  Future<void> registerWithEmailAndPassword(
    String email,
    String password,
  ) async {}
}
