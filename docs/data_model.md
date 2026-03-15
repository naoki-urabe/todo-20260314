# エンティティ
## User

| カラム | 型 | 制約　| 説明 | 
| --- | --- | --- | --- |
| id | UUID | PK | ユーザid。uuidを利用 |
| name | VARCHAR | NOT NULL | 表示ユーザ名 |
| email | VARCHAR | NOT NULL,UNIQUE | メールアドレス |
| is_active | BOOLEAN | NOT NULL | ユーザが有効かを表す　|
| password_hash | VARCHAR | NOT NULL| ユーザパスワードにsaltを付与しハッシュ化。argon2利用 |
| created_at | TIMESTAMP | NOT NULL | 作成日時 |
| updated_at | TIMESTAMP | NOT NULL | 公認日時 |

## Task

| カラム | 型 | 制約　| 説明 | 
| --- | --- | --- | --- |
| id | UUID | PK | タスクid。uuidを利用 |
| name | VARCHAR | NOT NULL |タスク名。 |
| user_id | VARCHAR | FOREIGN KEY (user_id) REFERENCES users(id) | ユーザid。外部キー|
| created_at | TIMESTAMP | NOT NULL | タスク作成日時 |
| updated_at | TIMESTAMP | NOT NULL | タスク更新日時 |
| deleted_at | TIMESTAMP | タスクの削除日時 |
| task_status | ENUM('TODO','IN_PROGRESS','DONE') | NOT NULL | タスクのステータス。TODO/IN_PROGRESS/DONE |