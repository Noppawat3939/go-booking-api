Table users {
id integer [primary key]
username text
role text [note: 'admin , null']
password text
active bool
created_at timestamp
updated_at timestamp
}

Table hotels {
id integer [primary key]
name text
description text
type text
price number
rooms_qty number
active bool
created_at timestamp
updated_up timestamp
}

Table hotel_facilities {
id integer [primary key]
hotel_id integer
wifi bool
break_fast bool
}

Table booking_hotels {
id integer [primary key]
hotel_id int
user_id int
status text [note: 'pending_payment, pending_checkin, checked_in, checked_out, cancelled']
days_qty number
check_in_at timestamp
check_out_at timestamp
}

Ref hotels_facilities: hotel_facilities.hotel_id > hotels.id // n-n
Ref: booking_hotels.hotel_id > hotels.id
Ref: booking_hotels.user_id > users.id
