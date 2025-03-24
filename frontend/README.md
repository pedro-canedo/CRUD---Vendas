# Sistema de Vendas - Frontend

Frontend do sistema de gerenciamento de vendas desenvolvido com Next.js e Material UI.

## Tecnologias Utilizadas

- Next.js 14
- React 18
- Material UI
- React Hook Form
- React Query
- React Toastify
- Date-fns
- Axios

## Pré-requisitos

- Node.js 18 ou superior
- npm ou yarn

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/seu-usuario/sistema-vendas.git
cd sistema-vendas/frontend
```

2. Instale as dependências:

```bash
npm install
# ou
yarn install
```

3. Crie um arquivo `.env.local` na raiz do projeto com as seguintes variáveis:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080/api
```

4. Inicie o servidor de desenvolvimento:

```bash
npm run dev
# ou
yarn dev
```

5. Acesse o sistema em `http://localhost:3000`

## Estrutura do Projeto

```
frontend/
├── src/
│   ├── app/              # Páginas e rotas
│   ├── components/       # Componentes reutilizáveis
│   ├── contexts/        # Contextos do React
│   ├── hooks/           # Hooks personalizados
│   ├── services/        # Serviços de API
│   └── utils/           # Funções utilitárias
├── public/              # Arquivos estáticos
└── package.json         # Dependências e scripts
```

## Funcionalidades

- Autenticação de usuários
- Dashboard com métricas
- Gestão de produtos
- Gestão de vendas
- Relatórios
- Configurações do sistema
- Backup do banco de dados
- Logs do sistema
- Perfil do usuário
- Tema escuro

## Scripts Disponíveis

- `npm run dev` - Inicia o servidor de desenvolvimento
- `npm run build` - Cria a build de produção
- `npm run start` - Inicia o servidor de produção
- `npm run lint` - Executa o linter
- `npm run format` - Formata o código com Prettier

## Contribuição

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
