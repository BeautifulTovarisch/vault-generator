use rouille::{ Request, Response };

pub fn vault(_req: &Request) -> Response {

    Response::text("Vault")
}
