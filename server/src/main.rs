use rouille::{ Response, router };

mod vault;
use vault::{ vault };

fn main() {
    rouille::start_server("0.0.0.0:3000", |request| {
        router!(request,
                (GET) (/) => {
                    Response::text("hello")
                },
                (POST) (/vault) => {
                    vault(request)
                },
                _ => Response::empty_404()
        )
    });
}
