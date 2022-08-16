package handlers

// Roots
const UserServiceRoot string = "http://localhost:8081/api/users/"
const DriveServiceRoot string = "http://localhost:8000/api/drives/"
const ReservationServiceRoot string = "http://localhost:8082/api/reservations/"
const ComplaintServiceRoot string = "http://localhost:8083/api/complaints/"
const RatingServiceRoot string = "http://localhost:8084/api/ratings/"

const _DriveServiceRoot string = "http://localhost:8000/api/drives"
const _ReservationServiceRoot string = "http://localhost:8082/api/reservations"
const _ComplaintServiceRoot string = "http://localhost:8083/api/complaints"
const _RatingServiceRoot string = "http://localhost:8084/api/ratings"

// Common
const Driver string = "driver/"
const Slash string = "/"

// UserService
const LoginApi = "login"
const AuthorizeApi string = "authorize/"
const _AuthenticateApi string = "authenticate"
const Registration string = "registration/"
const Search string = "search/"
const Admin string = "admin/"
const Passenger string = "passenger/"
const _Admin string = "admin"
const _Driver string = "driver"
const _Passenger string = "passenger"
const _Profile string = "profile"
const _ChangePassword string = "change-password"
const Ban string = "ban/"
const Verify string = "verify/"
const Unverified string = "unverified/"

// DriveService
const Finish string = "finish/"
const _Reserve string = "reserve"
const _Search string = "search"
const Finished string = "finished/"
const Unfinished string = "unfinished/"

// ReservationService
const _Verified = "verified"
const _Unverified = "unverified"
const User = "user/"
const Drive = "drive/"

// RatingService
const Evaluated = "evaluated/"

// Pageable
const pSearch = "search="
const pPage = "page="
const pSize = "size="
const QMark = "?"
const Amp = "&"
