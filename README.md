# Estudos

# 🛒 Aplicativo de Lista de Compras com Atualização em Tempo Real

## Descrição do Projeto
Este projeto é um **aplicativo colaborativo de lista de compras**, projetado para permitir que grupos de pessoas, como famílias ou amigos, compartilhem e atualizem listas em tempo real. Utilizando tecnologias modernas de backend e frontend, a aplicação oferece uma experiência dinâmica e eficiente para gerenciar as compras do dia a dia.

## Funcionalidades
- **Atualização em Tempo Real:** Utiliza WebSocket para garantir que qualquer mudança na lista seja instantaneamente refletida para todos os membros do grupo.
- **Sistema de Login:** Autenticação de usuários com JWT para garantir que somente pessoas autorizadas possam acessar as listas.
- **Gestão de Grupos:** Permite que grupos de usuários compartilhem e colaborem em múltiplas listas de compras.
- **Gerenciamento de Itens:** Adicione, edite e remova itens da lista com facilidade, além de marcar itens como comprados ou não comprados.
- **Responsividade:** A interface foi projetada para funcionar bem em diferentes dispositivos, como smartphones, tablets e desktops.

## Tecnologias Utilizadas
- **Frontend:** React para a construção da interface do usuário.
- **Backend:** Node.js com Express para o servidor e lógica de negócios.
- **Banco de Dados:** MySQL utilizando o Sequelize para o mapeamento objeto-relacional.
- **Autenticação:** JWT (JSON Web Token) para autenticação segura de usuários.
- **Comunicação em Tempo Real:** WebSocket para sincronização em tempo real entre os usuários.
- **Gerenciamento de Senhas:** Bcrypt para hash de senhas e segurança de dados.

## Como Executar o Projeto

### 1. Requisitos
- Node.js (v14 ou superior)
- MySQL (v8 ou superior)
- Yarn ou NPM

### 2. Configuração do Backend
\`\`\`bash
# Clone o repositório
git clone https://github.com/seu-usuario/nome-do-repositorio.git

# Instale as dependências
cd backend
yarn install

# Crie um arquivo .env com as seguintes variáveis:
# - DATABASE_URL: URL de conexão do MySQL
# - JWT_SECRET: Chave secreta para JWT
# - WEBSOCKET_PORT: Porta para o WebSocket

# Execute as migrações do banco de dados
yarn sequelize db:migrate

# Inicie o servidor
yarn start
\`\`\`

### 3. Configuração do Frontend
\`\`\`bash
# Instale as dependências
cd frontend
yarn install

# Inicie o frontend
yarn start
\`\`\`

### 4. Executar a Aplicação
Acesse a aplicação no navegador: \`http://localhost:3000\`.

## Estrutura do Projeto
\`\`\`bash
├── backend/                # Código do servidor e API
│   ├── models/             # Modelos do banco de dados (Sequelize)
│   ├── controllers/        # Lógica de controle
│   ├── routes/             # Rotas da API
│   ├── middlewares/        # Middlewares como autenticação JWT
│   └── sockets/            # Configuração do WebSocket
├── frontend/               # Código do frontend em React
│   ├── components/         # Componentes React reutilizáveis
│   ├── pages/              # Páginas da aplicação
│   ├── services/           # Serviços de API
│   └── hooks/              # Hooks customizados para gerenciar estado e lógica
└── README.md               # Documentação do projeto
\`\`\`

## Roadmap
- [x] Sistema de login e autenticação com JWT
- [x] Atualização em tempo real usando WebSocket
- [x] Gerenciamento de listas e itens
- [ ] Notificações de alteração de lista
- [ ] Implementação de lista offline com sincronização

## Contribuição
Contribuições são bem-vindas! Siga os passos abaixo:
1. Fork este repositório.
2. Crie uma nova branch (\`git checkout -b feature/nome-da-feature\`).
3. Commit suas alterações (\`git commit -m 'Adiciona nova feature'\`).
4. Envie para a branch original (\`git push origin feature/nome-da-feature\`).
5. Crie um Pull Request.

## Licença
Distribuído sob a licença MIT. Veja \`LICENSE\` para mais informações.