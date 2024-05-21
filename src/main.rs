use askama::Template;
use axum::extract::Query;
use axum::routing::get;
use axum::Router;
use serde::Deserialize;

#[derive(Deserialize)]
struct IndexQuery {
    /// External party
    p: Option<String>,
    /// Domain
    d: Option<String>,
}

#[derive(Template)]
#[template(path = "index.jinja2", ext = "html")]
struct IndexTemplate {
    result: String,
    external_party: String,
    domain: String,
}

async fn index(Query(input): Query<IndexQuery>) -> IndexTemplate {
    let result = if input.p.is_some() && input.d.is_some() {
        let party = input.p.as_ref().unwrap();
        let domain = input.d.as_ref().unwrap();
        Some(format!("ext.{party}.RANDOM@{domain}"))
    } else {
        None
    };

    IndexTemplate {
        result: result.unwrap_or_default(),
        external_party: input.p.unwrap_or_default(),
        domain: input.d.unwrap_or_default(),
    }
}

#[tokio::main]
async fn main() {
    let app = Router::new().route("/", get(index));

    // run it
    let listener = tokio::net::TcpListener::bind("127.0.0.1:3000")
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap();
}
