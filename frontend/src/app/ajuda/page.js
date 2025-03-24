'use client';

import {
    Box,
    Paper,
    Typography,
    Grid,
    Accordion,
    AccordionSummary,
    AccordionDetails,
    List,
    ListItem,
    ListItemText,
    ListItemIcon,
    Divider,
    Button,
} from '@mui/material';
import {
    ExpandMore as ExpandMoreIcon,
    QuestionAnswer as QuestionAnswerIcon,
    Email as EmailIcon,
    Phone as PhoneIcon,
    WhatsApp as WhatsAppIcon,
} from '@mui/icons-material';
import Navbar from '@/components/Navbar';

export default function Ajuda() {
    return (
        <Box>
            <Navbar />
            <Box
                component="main"
                sx={{
                    flexGrow: 1,
                    p: 3,
                    width: { sm: `calc(100% - 240px)` },
                    ml: { sm: '240px' },
                }}
            >
                <Typography variant="h4" gutterBottom>
                    Central de Ajuda
                </Typography>

                <Grid container spacing={3}>
                    <Grid item xs={12} md={8}>
                        <Paper sx={{ p: 3, mb: 3 }}>
                            <Typography variant="h6" gutterBottom>
                                Perguntas Frequentes
                            </Typography>

                            <Accordion>
                                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                                    <Typography>Como faço para cadastrar um novo produto?</Typography>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <Typography>
                                        Para cadastrar um novo produto, acesse o menu "Produtos" e
                                        clique no botão "Novo Produto". Preencha todos os campos
                                        obrigatórios e clique em "Salvar".
                                    </Typography>
                                </AccordionDetails>
                            </Accordion>

                            <Accordion>
                                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                                    <Typography>Como registro uma nova venda?</Typography>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <Typography>
                                        Para registrar uma nova venda, acesse o menu "Vendas" e
                                        clique no botão "Nova Venda". Selecione os produtos desejados,
                                        informe as quantidades e clique em "Finalizar Venda".
                                    </Typography>
                                </AccordionDetails>
                            </Accordion>

                            <Accordion>
                                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                                    <Typography>Como gero relatórios?</Typography>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <Typography>
                                        Para gerar relatórios, acesse o menu "Relatórios". Você pode
                                        filtrar por período e tipo de relatório desejado. Clique em
                                        "Gerar Relatório" para visualizar os dados.
                                    </Typography>
                                </AccordionDetails>
                            </Accordion>

                            <Accordion>
                                <AccordionSummary expandIcon={<ExpandMoreIcon />}>
                                    <Typography>Como faço backup do sistema?</Typography>
                                </AccordionSummary>
                                <AccordionDetails>
                                    <Typography>
                                        O sistema realiza backups automáticos conforme configurado.
                                        Para fazer um backup manual, acesse o menu "Backup" e clique
                                        em "Criar Backup Manual". Você também pode baixar backups
                                        anteriores.
                                    </Typography>
                                </AccordionDetails>
                            </Accordion>
                        </Paper>

                        <Paper sx={{ p: 3 }}>
                            <Typography variant="h6" gutterBottom>
                                Vídeos Tutoriais
                            </Typography>

                            <List>
                                <ListItem>
                                    <ListItemIcon>
                                        <QuestionAnswerIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="Como usar o sistema"
                                        secondary="Aprenda as funcionalidades básicas do sistema"
                                    />
                                </ListItem>
                                <ListItem>
                                    <ListItemIcon>
                                        <QuestionAnswerIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="Gestão de Produtos"
                                        secondary="Como cadastrar e gerenciar produtos"
                                    />
                                </ListItem>
                                <ListItem>
                                    <ListItemIcon>
                                        <QuestionAnswerIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="Gestão de Vendas"
                                        secondary="Como registrar e acompanhar vendas"
                                    />
                                </ListItem>
                                <ListItem>
                                    <ListItemIcon>
                                        <QuestionAnswerIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="Relatórios"
                                        secondary="Como gerar e analisar relatórios"
                                    />
                                </ListItem>
                            </List>
                        </Paper>
                    </Grid>

                    <Grid item xs={12} md={4}>
                        <Paper sx={{ p: 3 }}>
                            <Typography variant="h6" gutterBottom>
                                Suporte
                            </Typography>

                            <Typography variant="body2" paragraph>
                                Em caso de dúvidas ou problemas, entre em contato com nossa
                                equipe de suporte através dos canais abaixo:
                            </Typography>

                            <List>
                                <ListItem>
                                    <ListItemIcon>
                                        <EmailIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="E-mail"
                                        secondary="suporte@sistema.com"
                                    />
                                </ListItem>
                                <ListItem>
                                    <ListItemIcon>
                                        <PhoneIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="Telefone"
                                        secondary="(11) 1234-5678"
                                    />
                                </ListItem>
                                <ListItem>
                                    <ListItemIcon>
                                        <WhatsAppIcon />
                                    </ListItemIcon>
                                    <ListItemText
                                        primary="WhatsApp"
                                        secondary="(11) 98765-4321"
                                    />
                                </ListItem>
                            </List>

                            <Divider sx={{ my: 2 }} />

                            <Button
                                variant="contained"
                                fullWidth
                                startIcon={<EmailIcon />}
                                href="mailto:suporte@sistema.com"
                            >
                                Enviar E-mail
                            </Button>
                        </Paper>
                    </Grid>
                </Grid>
            </Box>
        </Box>
    );
} 