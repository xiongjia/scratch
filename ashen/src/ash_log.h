/*
 */

#ifndef _ASH_LOG_H_
#define _ASH_LOG_H_ 1

typedef struct _ash_log_context ash_log_context_t;


ash_log_context_t* ash_log_create(void);

void ash_log_destroy(ash_log_context_t *context);


#endif /* !defined(_ASH_LOG_H_) */

